using Formulator.Application.Interfaces;
using Formulator.Core.Entities;
using Formulator.Infrastructure.Data;
using System.ComponentModel.DataAnnotations;

namespace Formulator.Application.Services
{
    public class FormService : IFormService
    {
        private readonly ApplicationDbContext _context;

        public FormService(ApplicationDbContext context)
        {
            _context = context;
        }

        public Form CreateForm(Form form)
        {
            _context.Add(form);
            _context.SaveChanges();

            return form;
        }

        public Form GetForm(int id)
        {
            var form = _context.Forms.Find(id);
            ArgumentNullException.ThrowIfNull(form);
            return form;
        }
        public IEnumerable<Form> GetUserForms(User user)
        {
            var forms = _context.Forms.Where(f => f.CreatorId == user.Id).ToList();
            return forms;
        }

        public void UpdateForm(Form form)
        {
            var formEntity = _context.Forms.Find(form.Id);
            ArgumentNullException.ThrowIfNull(formEntity);
            _context.Entry(formEntity).CurrentValues.SetValues(form);
        }

        public void DeleteForm(int id)
        {
            var formEntity = _context.Forms.Find(id);
            ArgumentNullException.ThrowIfNull(formEntity);
            _context.Remove(formEntity);
        }
    }
}